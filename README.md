# README

Concatenate piped pdf files.


```PowerShell
(ls *pdf).FullName|.\go-pdfconc.exe --outname hoge
```

---

Using [pdfcpu](pdfcpu.io) (Apache-2.0 license).

```
go get github.com/pdfcpu/pdfcpu/cmd/...
```
