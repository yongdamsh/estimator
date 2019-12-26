package task

const Template = `
<!DOCTYPE html>
<html>
  <head>
    <title>Task Estimator</title>
  </head>
  <body>
    <h1>Tasks</h1>
    
		<ol>
			{{range .}}
				<li>{{.Name}} ({{.Elapsed}}/{{.CurEst}})</li>
			{{end}}
		</ol>

		<form name="task" method="POST" action="/tasks">
			<label>Name: <input type="text" name="name" /></label>
		</form>
  </body>
</html>
`
