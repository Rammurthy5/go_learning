module example.com/hello

go 1.16

replace example.com/greetings => ../greetings

require (
	example.com/err_greeting v0.0.0-00010101000000-000000000000
	example.com/greetings v0.0.0-00010101000000-000000000000
	example.com/slices_learning v0.0.0-00010101000000-000000000000
)

replace example.com/err_greeting => ../err_greeting

replace example.com/slices_learning => ../slices_learning
