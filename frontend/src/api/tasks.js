export function getAll() {
  return fetch('http://localhost:8000/tasks');
}

export function store(data) {
  return fetch('http://localhost:8000/tasks', {
    method: 'POST',
    body: JSON.stringify(data),
  })
}

export function complete(id) {
  return fetch('http://localhost:8000/tasks/' + id + '/complete', {
    method: 'PUT',
  })
}