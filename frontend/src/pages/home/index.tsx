import './index.css'
import { mutate }  from 'swr';
import AppendTodo from '../../components/AppendTodo';
import { useEffect, useState } from 'react';
import axios from 'axios';
import { CheckCircleFillIcon, CircleIcon } from '@primer/octicons-react';

export interface Todo {
  id: number
  title: string
  completed: boolean
  description: string
}

export const ENDPOINT = "http://localhost:8000";

    
export default function Home(props: {name: string, id:string}) {

    const [reports, setReports] = useState([]);
    
    useEffect(() => {
      axios.get(`${ENDPOINT}/api/todos`).then(res => {
        setReports(res.data)
      }).catch(err => console.log("error: ", err))
    }), [reports];
    
    async function completeTodo(id: number) {
      await axios.patch(`${ENDPOINT}/api/todos/${id}/completed`).then(res => {
        setReports(res.data)
      }).catch(err => console.log(err));
    }
  
    async function undoTodo(id: number) {
      await axios.patch(`${ENDPOINT}/api/todos/${id}/uncompleted`).then(res => {
        setReports(res.data)
      }).catch(err => console.log("ERR:", err));
    }
    return (
            <>
            <h1 className='text-light'>{props.name ? "Welcome, " + props.name + '!': 'You are not authorized'}</h1>
            <table className='table table-striped table-dark'>
                        <thead>
                    <tr>
                        <th>Id</th>
                        <th>Title</th>
                        <th>Description</th>
                        <th>Completed</th>
                    </tr>
                        </thead>
                        <tbody>
                        {reports?.map ((todo:Todo) => {
                        return (
                            <tr key={todo.id}>
                            <td>{todo.id}</td>
                            <td>{todo.title}</td>
                            <td>{todo.description}</td>
                            {todo.completed === true ? (<td onClick={() => undoTodo(todo.id)}><CheckCircleFillIcon/></td>) : (<td onClick={() => completeTodo(todo.id)}> <CircleIcon/></td>)}
                            </tr>
                        )
                        })}      
                        </tbody>
                    </table> 

                <AppendTodo mutate={mutate} id={props.id}/>
            </>
    )
}