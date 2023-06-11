import { useState } from "react";
import { Button, Group, Modal } from "@mantine/core";
import { ENDPOINT, Todo } from "../App";
import { KeyedMutator } from "swr";
import axios from "axios";

function AppendTodo({mutate}: {mutate : KeyedMutator<Todo[]>}) {
    const [open, setOpen] = useState(false);
    let initialState: Todo = {
        id: 0,
        title: "",
        completed: false,
        description: ""
    }
    const [report, setReport] = useState<Todo>(initialState)

    const submitForm = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        axios.post(`${ENDPOINT}/api/todos`, report).then((res) => {
            setReport(res.data);
            console.log(res.data);
        }).catch(err => console.log("ERR", err));
    }

    const onChangeHandler = (event: HTMLInputElement) => {
        const {name, value} = event
        setReport((prev) => {
            return {...prev, [name]: value}
        })
    }

    return (
        <>
            <Modal opened={open} onClose={() => setOpen(false)} title="Create new Report!">
                <form onSubmit={submitForm}>
                    <label>Title</label>
                    <input className="form-control" name="title" value={report.title} onChange={(e) => onChangeHandler(e.target)}/>
                    <label>Description</label>
                    <input className="form-control" name="description" value={report.description} onChange={(e) => onChangeHandler(e.target)}/>

                    <button type="submit" className="btn btn-primary mt-2">Create</button>
                </form>
            </Modal>        
            <Group position="center">
                <Button fullWidth mb={12} onClick={() => setOpen(true)}>
                    ADD
                </Button>
            </Group>
        </>
    )
}

export default AppendTodo;