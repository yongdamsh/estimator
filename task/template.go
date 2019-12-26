package task

const (
	ListTemplate = `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Task Estimator</title>
			</head>
			<body>
				<h1>Tasks</h1>

				<a href="/tasks/new">Create</a>
				<hr />
				
				<ol>
					{{range .Tasks}}
						<li>{{.Name}} ({{.Elapsed}}/{{.CurEst}})</li>
					{{end}}
				</ol>
			</body>
		</html>
	`
	NewTaskTemplate = `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Task Estimator</title>
			</head>
			<body>
				<h1>Create a New Task</h1>
				
				<form name="task" method="post" action="/tasks">
					<label>Name: <input type="text" name="name" /></label>
					<br />
					<label>Estimated Time: <input type="text" name="estimatedtime" placeholder="1h 30m" /></label>
					<br />
					<button type="submit">Create</button>
				</form>
			</body>
		</html>
	`
)
