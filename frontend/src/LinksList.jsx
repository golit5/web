import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const LinksList = () => {
    const [links, setLinks] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        fetch('http://localhost:8080/links',)
            .then(response => {
                console.log(response)
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            }) 
            .then(data => {
                setLinks(data); 
            })
            .catch(error => console.log(error));
    }, []);

    const handleDelete = (id) => {
        fetch(`http://localhost:8080/links/${id}`, {
            method: 'DELETE',  
        })
        .then(response => {
            console.log(response)
            if (response.ok) {
                setLinks(links.filter(link => link.id !== id)); // Убираем удаленную ссылку из списка
            }
        })
        .catch(error => console.log(error));
    };

    const handleUpdate = (id) => {
        navigate(`/update/${id}`);
    };

    return (
        <div className="general" id="links">
            <h2>Links</h2>
            <ul>
                {links.map(link => (
                    <li key={link.id}>
                        <a href={link.url} target="_blank" rel="noreferrer">{link.title}</a>
                        <div>
                        {/* <button class="submit">
                            <Link to={`/update/${link.id}`}>Update</Link>
                        </button> */}
                        <button class="submit" onClick={() => handleUpdate(link.id)}>Update</button>
                        <button class="delete" onClick={() => handleDelete(link.id)}>Delete</button>
                        </div>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default LinksList;
