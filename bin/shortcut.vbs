Set oWS = WScript.CreateObject("WScript.Shell") 
sLinkFile = oWS.ExpandEnvironmentStrings("%fld%\%title%.lnk")
Set oLink = oWS.CreateShortcut(sLinkFile) 
oLink.TargetPath = oWS.ExpandEnvironmentStrings("%windir%\system32\cmd.exe")
oLink.Arguments = oWS.ExpandEnvironmentStrings(" /c start """" /b bin\shell\bash -c ""disposable %action%""")
oLink.WindowStyle = oWS.ExpandEnvironmentStrings("%winStyle%")
oLink.IconLocation = oWS.ExpandEnvironmentStrings("%icon%")
oLink.WorkingDirectory = oWS.ExpandEnvironmentStrings("%dir%")
oLink.Description = oWS.ExpandEnvironmentStrings("%action% %vm% disposable VM")
oLink.Save