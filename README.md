[![Go Reference](https://pkg.go.dev/badge/github.com/kjk/windigo.svg)](https://pkg.go.dev/github.com/kjk/windigo)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/rodrigocfd/windigo.svg)](https://github.com/kjk/windigo)
[![Lines of code](https://tokei.rs/b1/github/rodrigocfd/windigo)](https://github.com/kjk/windigo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Windigo

Win32 API and GUI in idiomatic Go.

## Overview

The UI library is divided in the following packages:

| Package | Description |
| - | - |
| `ui` | High-level UI wrappers for windows and controls. |
| `ui/wm` | High-level event parameters (Windows message callbacks). |

For the Win32 API bindings:

| Package | Description |
| - | - |
| [`win`](win/) | Native Win32 structs, handles and functions. |
| [`win/co`](win/co/) | Native Win32 constants, all typed. |
| `win/errco` | Native Win32 [error codes](https://docs.microsoft.com/en-us/windows/win32/debug/system-error-codes), with types `errco.ERROR` and `errco.CDERR`. |

For the [COM](win/com/) bindings, there is the main package, and two subpackages – the `co` suffix contains the constants, and the `vt` contains the virtual tables:

| Packages | Description |
| - | - |
| `win/com/autom`<br>`win/com/autom/automco`<br>`win/com/autom/automvt` | Native Win32 [Automation](https://docs.microsoft.com/en-us/windows/win32/api/_automat/) COM interfaces. |
| `win/com/com`<br>`win/com/com/comco`<br>`win/com/com/comvt` | Native Win32 [COM API base](https://docs.microsoft.com/en-us/windows/win32/api/_com/). |
| `win/com/d2d1`<br>`win/com/d2d1/d2d1co`<br>`win/com/d2d1/d2d1vt` | Native Win32 [Direct2D](https://docs.microsoft.com/en-us/windows/win32/direct2d/direct2d-portal) COM interfaces. |
| `win/com/dshow`<br>`win/com/dshow/dshowco`<br>`win/com/dshow/dshowvt` | Native Win32 [DirectShow](https://docs.microsoft.com/en-us/windows/win32/directshow/directshow) COM interfaces. |
| `win/com/shell`<br>`win/com/shell/shellco`<br>`win/com/shell/shellvt` | Native Win32 [Shell](https://docs.microsoft.com/en-us/windows/win32/api/_shell/) COM interfaces. |

Windigo is designed to be familiar to Win32 programmers, using the same concepts, so most C/C++ Win32 tutorials should be applicable.

Windows and controls can be created in two ways:

* programmatically, by specifying the options used in the underlying [CreateWindowEx](https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw);
* by loading [resources](https://en.wikipedia.org/wiki/Resource_(Windows)#Resource_software) from a `.rc` or a `.res` file.

CGo is **not** used, just syscalls.

## Error treatment

The native Win32 functions deal with errors in two ways:

* Recoverable errors will return an `errco.ERROR` value, which implements the [`error`](https://golang.org/pkg/builtin/#error) interface;

* Unrecoverable errors will simply **panic**. This avoids the excess of `if err != nil` with errors that cannot be recovered anyway, like internal Windows errors.

## Example

The example below creates a window programmatically, and handles the button click. Also, it uses the `minimal.syso` provided in the [resources](resources/) folder.

![Screen capture](example.gif)

```go
package main

import (
    "fmt"
    "runtime"

    "github.com/kjk/windigo/ui"
    "github.com/kjk/windigo/win"
    "github.com/kjk/windigo/win/co"
)

func main() {
    runtime.LockOSThread()

    myWindow := NewMyWindow() // instantiate
    myWindow.wnd.RunAsMain()  // ...and run
}

// This struct represents our main window.
type MyWindow struct {
    wnd     ui.WindowMain
    lblName ui.Static
    txtName ui.Edit
    btnShow ui.Button
}

// Creates a new instance of our main window.
func NewMyWindow() *MyWindow {
    wnd := ui.NewWindowMain(
        ui.WindowMainOpts().
            Title("Hello you").
            ClientArea(win.SIZE{Cx: 340, Cy: 80}).
            IconId(101), // ID of icon resource, see resources folder
    )

    me := &MyWindow{
        wnd: wnd,
        lblName: ui.NewStatic(wnd,
            ui.StaticOpts().
                Text("Your name").
                Position(win.POINT{X: 10, Y: 22}),
        ),
        txtName: ui.NewEdit(wnd,
            ui.EditOpts().
                Position(win.POINT{X: 80, Y: 20}).
                Size(win.SIZE{Cx: 150}),
        ),
        btnShow: ui.NewButton(wnd,
            ui.ButtonOpts().
                Text("&Show").
                Position(win.POINT{X: 240, Y: 19}),
        ),
    }

    me.btnShow.On().BnClicked(func() {
        msg := fmt.Sprintf("Hello, %s!", me.txtName.Text())
        me.wnd.Hwnd().MessageBox(msg, "Saying hello", co.MB_ICONINFORMATION)
    })

    return me
}
```

## License

Licensed under [MIT license](https://opensource.org/licenses/MIT), see [LICENSE.md](LICENSE.md) for details.
