import { Link, useNavigate } from "react-router-dom"
import "../../pages/register/index.css"
import { SyntheticEvent, useState } from "react"
import { ENDPOINT } from "../../App";
export default function Login(props: {setName: (name:string) => void}) {

    const navigate = useNavigate()

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');


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
        if(result.name) {
            props.setName(result.name)
            return navigate("/home");
        } else {
            setError("Invalid login")
        }
    }
    return(
        <>
        <div className="wrapper fadeInDown">
            <div id="formContent">
                <form onSubmit={handleSubmit}>
                    <input type="text" id="email" className="fadeIn second" onChange={e => setEmail(e.target.value)} name="email" placeholder="E-mail"/>
                    <input type="password" id="password" className="fadeIn third" name="password" onChange={e => setPassword(e.target.value)} placeholder="Password"/>
                    <input type="submit" className="fadeIn fourth" value="Log In"/>
                </form>
                <div className="text-danger">
                    <h4>{error}</h4>
                </div>
                <div id="formFooter">
                <Link to="/register" className="underlineHover">Don’t have an account? ?</Link>
                </div>
            </div>
        </div>
        </>

    )
}