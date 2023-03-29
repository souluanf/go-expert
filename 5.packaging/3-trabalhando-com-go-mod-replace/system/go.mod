module github.com/souluanf/go-expert/5.packaging/3-trabalhando-com-go-mod-replace/system

go 1.20

replace github.com/souluanf/go-expert/5.packaging/3-trabalhando-com-go-mod-replace/math => ../math

require github.com/souluanf/go-expert/5.packaging/3-trabalhando-com-go-mod-replace/math v0.0.0-00010101000000-000000000000
