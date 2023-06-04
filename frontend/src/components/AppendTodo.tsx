import { useState } from "react";
import  {useForm}  from '@mantine/form';
import { Button, Group, Modal, TextInput, Textarea } from "@mantine/core";
import { ENDPOINT, Todo } from "../App";
import { KeyedMutator } from "swr";

function AppendTodo({mutate}: {mutate : KeyedMutator<Todo[]>}) {
    const [open, setOpen] = useState(false);

    const form = useForm({
        initialValues: {
            title: "",
            description:"",
        },
    })

    async function createTodo(values: {title: string, description: string}) {
        const updated = await fetch(`${ENDPOINT}/api/todos`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(values)
        }).then((r) => r.json());

        mutate(updated)
        form.reset()
        setOpen(false)
    }

    return (
        <>
            <Modal opened={open} onClose={() => setOpen(false)} title="create new Todo">
                <form onSubmit={form.onSubmit(createTodo)}>
                    <TextInput required mb={12} label="Todo" placeholder="Aqui" {...form.getInputProps("title")}/>
                    <Textarea  required mb={12} label="Description" placeholder="Aqui" {...form.getInputProps("description")}/>

                    <Button type="submit">Create</Button>
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