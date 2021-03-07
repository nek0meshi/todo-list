import http from '../utils/http'

export function getAll() {
  return http.get(baseUrl())
}

export function store(data) {
  return http.post(baseUrl(), JSON.stringify(data))
}

export function complete(id) {
  return http.put(baseUrl() + '/' + id + '/complete')
}

export function update(id, data) {
  return http.put(baseUrl() + '/' + id + '/update', JSON.stringify(data))
}

function baseUrl() {
  return 'http://localhost:8000/tasks';
}