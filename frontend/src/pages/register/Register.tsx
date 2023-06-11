import { Link } from "react-router-dom"
import "./index.css"
import { SyntheticEvent, useState } from "react"
import { ENDPOINT } from "../home";
import { useNavigate } from "react-router-dom";
export default function Register() {
    let navigate = useNavigate();

    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        await fetch(`${ENDPOINT}/api/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }, 
            body: JSON.stringify({
                name,
                email,
                password
            }),
        }).then((res) => {console.log("RES:",res)});
        return navigate("/login");
    }
    
    
    return(
        <>
        <div className="wrapper fadeInDown">
            <div id="formContent">
                <div className="fadeIn first">
                </div>

                <form onSubmit={submit}>
                    <input type="text" id="name"  onChange={e => setName(e.target.value)} className="fadeIn second" name="name" placeholder="Name"/>
                    <input type="text" id="email" onChange={e => setEmail(e.target.value)} className="fadeIn second" name="email" placeholder="Email"/>
                    <input type="password" id="password" onChange={e => setPassword(e.target.value)} className="fadeIn third" name="password" placeholder="Password"/>
                    <input type="submit" className="fadeIn fourth" value="Register"/>
                </form>
                <div id="formFooter">
                <Link to="/login" className="underlineHover">Has an account?</Link>
                </div>
            </div>
        </div>
        </>

    )
}