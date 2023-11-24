import { useState, SyntheticEvent } from "react";
import { Button, Group, Modal } from "@mantine/core";
import { ENDPOINT } from "../App";

function AppendTodo(props: {id:string}) {
    const [open, setOpen] = useState(false);
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [completed, setCompleted] = useState("false");

    async function handleSubmit(e: SyntheticEvent) {
        const user_id = (props.id).toString();
        setCompleted("false");
        e.preventDefault();
        const res = await fetch(`${ENDPOINT}/api/data/create`, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                title,
                description,
                completed,
                user_id
            })
        })
        console.log("RES:", res)

    }

    return (
        <>
            <Modal opened={open} onClose={() => setOpen(false)} title="Create new Report!">
                <form onSubmit={handleSubmit}>
                    <label>Title</label>
                    <input className="form-control" name="title" value={title} onChange={(e) => setTitle(e.target.value)}/>
                    <label>Description</label>
                    <input className="form-control" name="description" value={description} onChange={(e) => setDescription(e.target.value)}/>
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