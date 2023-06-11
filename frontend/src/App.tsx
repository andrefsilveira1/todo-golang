import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import Nav from './components/navbar'
import Login from './components/login'
import Home from './pages/home'
export const ENDPOINT = "http://localhost:8000";
import 'bootstrap/dist/css/bootstrap.css';
import Register from './pages/register/Register'

export interface Todo {
  id: number
  title: string
  completed: boolean
  description: string
}
function App() {
  

  return (
    <>
      <BrowserRouter>
      <Nav/>
        <Routes>
          <Route path="/login" Component={Login} />
          <Route path="/register" Component={Register} />
          <Route path="/home" Component={Home} />
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
