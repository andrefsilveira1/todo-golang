import { Link, useNavigate } from "react-router-dom"
import "./index.css"
import { SyntheticEvent, useState } from "react"
import { ENDPOINT } from "../../App";
export default function Login(props: {setName: (name:string) => void}) {

    const navigate = useNavigate()

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    async function handleSubmit(e: SyntheticEvent) {
        e.preventDefault();
        const res = await fetch(`${ENDPOINT}/api/login`, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        })
        const result = await res.json();
        console.log("RESULT:", result)
        props.setName(result.name)
        return navigate("/home");
    }
    return(
        <>
        <div className="wrapper fadeInDown">
            <div id="formContent">

                <div className="fadeIn first">
                </div>

                <form onSubmit={handleSubmit}>
                    <input type="text" id="email" className="fadeIn second" onChange={e => setEmail(e.target.value)} name="email" placeholder="email"/>
                    <input type="text" id="password" className="fadeIn third" name="password" onChange={e => setPassword(e.target.value)} placeholder="Password"/>
                    <input type="submit" className="fadeIn fourth" value="Log In"/>
                </form>
                <div id="formFooter">
                <Link to="/register" className="underlineHover">Don’t have an account? ?</Link>
                </div>
            </div>
        </div>
        </>

    )
}