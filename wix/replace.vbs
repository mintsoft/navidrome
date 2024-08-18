Const ForReading = 1    
Const ForWriting = 2

sFilename = Wscript.Arguments(0)

Set oFSO = CreateObject("Scripting.FileSystemObject")
Set oFile = oFSO.OpenTextFile(sFilename, ForReading)
sFileContent = oFile.ReadAll
oFile.Close

sNewFileContent = Replace(sFileContent, "[MSI_PLACEHOLDER_SECTION]" & vbCrLf, "")
Set oFile = oFSO.OpenTextFile(sFilename, ForWriting)
oFile.Write sNewFileContent
oFile.Close
