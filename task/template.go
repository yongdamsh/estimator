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
					{{range .}}
						<li>
							<p>{{.Name}} ({{.Elapsed}}/{{.Estimated}}) <a href="/tasks/edit/{{.Id}}">Edit</a></p>
						</li>
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
				
				<form name="task" method="post" action="/tasks/">
					<select name="featureId" required>
						<option value="">Choose a feature</option>
						{{range .}}
							<option value={{.Id}}>{{.Name}}</option>
						{{end}}
					</select>
					<label>Name: <input type="text" name="name" required /></label>
					<br />
					<label>Estimated Time: <input type="text" name="estimated" placeholder="1h 30m" required /></label>
					<br />
					<button type="submit">Create</button>
				</form>
			</body>
		</html>
	`
	EditTaskTemplate = `
	<!DOCTYPE html>
		<html>
			<head>
				<title>Task Estimator</title>
			</head>
			<body>
				<h1>Edit Task ID {{.Task.Id}}</h1>
				
				<form name="task" method="post" action="/tasks/edit/{{.Task.Id}}">
					<select name="featureId" required>
						<option value="">Choose a feature</option>
						{{$featureId := .Task.Feature.Id}}
						{{range .Features}}
							{{if eq .Id $featureId}}
								<option value={{.Id}} selected>{{.Name}}</option>
							{{else}}
								<option value={{.Id}}>{{.Name}}</option>
							{{end}}
						{{end}}
					</select>
					<label>Name: <input type="text" name="name" required value={{.Task.Name}} /></label>
					<br />
					<label>Estimated Time: <input type="text" name="estimated" placeholder="1h 30m" required value={{.Task.Estimated}} /></label>
					<br />
					<label>Elapsed Time: <input type="text" name="elapsed" placeholder="1h 30m" value={{.Task.Elapsed}} /></label>
					<br />
					<button type="submit">Edit</button>
				</form>
			</body>
		</html>
	`
)
