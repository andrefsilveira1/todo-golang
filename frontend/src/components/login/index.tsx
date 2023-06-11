import { Link } from "react-router-dom"
import "./index.css"
export default function Login() {
    return(
        <>
        <div className="wrapper fadeInDown">
            <div id="formContent">

                <div className="fadeIn first">
                <img src="http://danielzawadzki.com/codepen/01/icon.svg" id="icon" alt="User Icon" />
                </div>

                <form>
                    <input type="text" id="login" className="fadeIn second" name="login" placeholder="Login"/>
                    <input type="text" id="password" className="fadeIn third" name="login" placeholder="Password"/>
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