@echo off
set fld=%~2\Disposable %~1
set vm=%~1
set dir=%CD%
set winStyle=7
echo.
echo Creating %fld%...
echo.
mkdir "%fld%"
set title=Start %~1 Disposable
set action=start
set icon=%SystemRoot%\system32\SHELL32.dll, 299
cscript //nologo bin\shortcut.vbs
set title=Dispose %~1 Disposable
set action=dispose
set icon=%SystemRoot%\system32\SHELL32.dll, 219
cscript //nologo bin\shortcut.vbs
set title=Edit %~1 Permanent
set winStyle=1
set action=edit
set icon=%SystemRoot%\system32\SHELL32.dll, 316
cscript //nologo bin\shortcut.vbs
