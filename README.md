### Installation

```go
git clone github.com/syntaxgame/tbledit
cd tbledit
go build -o tbledit
```
or
<a href="https://github.com/syntaxgame/tbledit/releases"> Download a Release</a>.


### How to Use?
You can convert a .tbl file into excel table as following:
```
./tbledit export -i [tbl_file] -o [xlsx_file]
```

After tbl file is exported, you can edit the table on Excel, LibreOffice or Google Drive Sheets.

Finally, you can convert an excel table into .tbl file as following:
```
./tbledit import -i [xlsx_file] -o [tbl_file]
```

### Example
```
./tbledit export -i tb_cashshop.tbl -o cashshop.xlsx
./tbledit import -i cashshop.xlsx -o tb_cashshop.tbl
```
