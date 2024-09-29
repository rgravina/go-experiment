import {useEffect, useState} from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

interface User {
    name: String
}

function UserDetails({user}: { user: User }) {
    return <>
        <span>User is: {user.name}</span>
    </>;
}

function App() {
    const [count, setCount] = useState(0)

    const [user, setUser] = useState<User | undefined>(undefined);

    useEffect(() => {
        const fetchData = async () => {
            const response = await fetch("http://localhost:8080/api");
            const user = await response.json();
            setUser(user);
        };

        fetchData().catch((err) => console.log(err));
    }, []);


    return (
        <>
            <div>
                <a href="https://vitejs.dev" target="_blank">
                    <img src={viteLogo} className="logo" alt="Vite logo"/>
                </a>
                <a href="https://react.dev" target="_blank">
                    <img src={reactLogo} className="logo react" alt="React logo"/>
                </a>
            </div>
            <h1>Vite + React</h1>
            {user && <UserDetails user={user}/>}
            <div className="card">
                <button onClick={() => setCount((count) => count + 1)}>
                    count is {count}
                </button>
                <p>
                    Edit <code>src/App.tsx</code> and save to test HMR
                </p>
            </div>
            <p className="read-the-docs">
                Click on the Vite and React logos to learn more
            </p>
        </>
    )
}

export default App
