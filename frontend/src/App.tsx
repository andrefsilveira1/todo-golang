import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import Nav from './components/navbar'
import Login from './components/login'
import Home from './pages/home'
export const ENDPOINT = "http://localhost:8000";
import 'bootstrap/dist/css/bootstrap.css';
import Register from './pages/register/Register'
import { useEffect, useState } from 'react'

export interface Todo {
  id: number
  title: string
  completed: boolean
  description: string
}
function App() {

  const [name, setName] = useState('');
  const [id, setId] = useState('');
  useEffect(() => {
    (
      async () => {
        const res = await fetch(`${ENDPOINT}/api/user`, {
          headers: {'Content-Type': 'application/json'},
          credentials: 'include',
      });
        const result = await res.json();
        setName(result.name);
        setId(result.id);
      }
    )();
  });
  

  return (
    <>
      <BrowserRouter>
      <Nav name={name} setName={setName}/>
        <Routes>
          <Route path="/login" Component={() => <Login setName={setName}/>} />
          <Route path="/register" Component={Register} />
          <Route path="/home" Component={() => <Home name={name} id={id} />} />
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
