import { Box, List, ThemeIcon } from '@mantine/core'
import './App.css'
import  useSWR  from 'swr';
import AppendTodo from './components/AppendTodo';
import { CheckCircleFillIcon } from '@primer/octicons-react';

export interface Todo {
  id: number
  title: string
  completed: boolean
  description: string
}

export const ENDPOINT = "http://localhost:8000";

const fetcher = (url:string) => fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

function App() {
  const { data, mutate } = useSWR<Todo[]>('api/todos', fetcher)

  async function completeTodo(id: number) {
    const updated = await fetch(`${ENDPOINT}/api/todos/${id}/completed`, {
      method: "PATCH",

    }).then((r) => r.json());
    mutate(updated);
  }
  return (
    <>
    <Box
      sx={(theme) => ({
        padding: "2rem",
        width: "100%",
        maxWidth: "40rem",
        margin: "0 auto"
      })}
    >
      <List spacing="xs" size="sm" mb={12} center >
        {data?.map ((todo) => {
          return (
            <List.Item key={`todo_${todo.id}`} onClick={() => completeTodo(todo.id)}>
            icon={
              todo.completed ? (<ThemeIcon color='teal' size={24} radius="xl"> <CheckCircleFillIcon size={20}/></ThemeIcon>) : (<ThemeIcon color='gray' size={24} radius="xl"> <CheckCircleFillIcon size={20}/></ThemeIcon>)
            }
              {todo.title}
            </List.Item>
          )
        })}      
      </List>
      <AppendTodo mutate={mutate}/>
    </Box>
    </>
  )
}

export default App
