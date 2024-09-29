import {useEffect, useState} from 'react'
import './App.css'

interface User {
    name: String
}

function UserDetails({users}: { users: [User] }) {
    return <div className="user-details">
        <table>
            <thead>
            <tr>
                <th>Name</th>
            </tr>
            </thead>
            <tbody>
            {users?.map((user) => {
                return <tr>
                    <td>{user.name}</td>
                </tr>
            })
            }
            </tbody>
        </table>
    </div>
}

function App() {
    const [users, setUsers] = useState<[User] | undefined>(undefined);

    useEffect(() => {
        const fetchData = async () => {
            const response = await fetch("http://localhost:8080/api/users");
            const users = await response.json();
            setUsers(users);
        };

        fetchData().catch((err) => console.log(err));
    }, []);


    return (
        <>
            <h1>React app served by Echo</h1>
            <p>Here is a list of names stored in a SQLite database table:</p>
            <div className="card">
                {users && <UserDetails users={users}/>}
            </div>
        </>
    )
}

export default App
