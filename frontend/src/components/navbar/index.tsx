import { Link } from "react-router-dom";
import { ENDPOINT } from "../../App";

export default function Nav(props: {name:string, setName: (name:string) => void}) {

  let menu;
  async function logout () {
    await fetch(`${ENDPOINT}/api/logout`, {
      method:'POST',
      headers: {'Content-Type': 'application/json'},
      credentials: 'include'
    });
    props.setName('');
  }
  if(props.name === '') {
    menu = (
      <ul className="navbar-nav">
        <li className="nav-item active">
          <Link to="/home" className="nav-link">Home</Link>
        </li>
        <li className="nav-item">
          <Link  to="/login" className="nav-link" >Login</Link>
        </li>
        <li className="nav-item">
          <Link to="/register" className="nav-link">Register</Link>
        </li>
      </ul>
    )
  } else {
    menu = (
    <ul className="navbar-nav me-auto">
      <li className="nav-item active">
        <Link to="/login" className="nav-link" onClick={logout}>Logout</Link>
      </li>
    </ul>
      )
  }
    return (
      <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
        <div className="container-fluid">
          <a className="navbar-brand" href="#">Navbar</a>
          <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarNavDropdown">
            {menu}
          </div>
        </div>
      </nav>
    )
}