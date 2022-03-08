const path = require('path');
const app = require('express')();
const bodyParser = require('body-parser');
const notifier = require('node-notifier');
const port = process.env.PORT || 9000;
app.use(bodyParser.json());

app.get('/health', (req, res) => {
  res.status(200).send();
});

app.post('/notify', (req, res) => {
  notify(req.body, reply => res.send(reply))
});

app.listen(port, () => {
  console.log(`Server listening on port ${port}`);
});

//notify function
const notify = ({ title, message }, cb) => {
  notifier.notify(
    {
      title: title || 'Notification',
      message: message || 'Notification message',
      sound: true,
      wait: true,
      reply: true,
      closeLabel: 'Finish Task',
      timeout: 15,
      icon: path.join(__dirname, '/assets/apple.png')
    },
    (err, response, reply) => {
      cb(reply);
    }
  )
};