package generator

import (
	"fmt"
	"log"
	"mdx/ast"
	"os"
)

func GenerateDocument(filename string, elements []ast.Component) {
	file, fileErr := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if fileErr != nil {
		log.Fatal(fileErr.Error())
	}

	defer file.Close()

	file.WriteString(`<html>
    <head>
        <script>
            const handleClick = () => {
                console.log("Hello!");
            }
        </script>
        <style>
        	*, *::before, *::after {
        		box-sizing: border-box;
        		padding: 0;
        		margin: 0;
	        	font-family: sans-serif;
        	}

        	body {
        		margin: 24px;
        	}

        	code {
        		font-family: monospace;
        	}

        	blockquote {
	        	margin: 0;
			    padding: 0 1em;
			    color: #57606a;
			    border-left: .25em solid #d0d7de;
		    }

		    .code-block {
			    display: block;
			    border-radius: 4px;
			    background-color: #1c373d;
			    color: #dbdbe3;
			    counter-reset: line;
			    position: relative;
			    margin: 16px 0;
			    overflow-x: auto;
			    scrollbar-width: thin;
		    }

		    .code-block > pre {
			    font-family: 'Consolas', monospace;
		    }

			.code-block > pre::before {
			    counter-increment: line;
			    content: counter(line);
			    display: inline-block;
			    width: 40px;
			    background-color: #030c0e;
			    padding: 2px 8px 2px 2px;
			    margin-right: 12px;
			    text-align: right;
			    color: #7f7a7a;
			}

			.primary-btn {
				cursor: pointer;
				border: none;
				margin: none;
				padding: 8px 12px;
				background-color: teal;
				color: whitesmoke;
				border-radius: 4px;
			}
        </style>
    </head>
`)

	body := &ast.Body{Children: elements}
	n, writeErr := file.WriteString(body.Html())
	if writeErr != nil {
		log.Fatal(writeErr.Error())
	}

	file.WriteString(`
</html>
`)

	fmt.Printf("%d autogenerated bytes written to %s\n", n, filename)
}
