const express = require('express')
const app = express()
const port = 3000

app.use(express.static('public'))
// Parse URL-encoded bodies (as sent by HTML forms)
app.use(express.urlencoded({ extended: true }));
// Parse JSON bodies (as sent by API clients)
app.use(express.json());


app.get('/data', async (req, res) => {
  const re = await fetch('https://jsonplaceholder.typicode.com/todos/')
  const r = await re.json()
  res.send(r)
})

app.post('/add', (req, res) => {
    const name = req.body.name
    const email = req.body.email
    console.log(req.body)
    res.send(`
        <div>
            <p>${name}</p>
            <p>${email}</p>
        </div>
    `)
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})