import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const AddLink = () => {
    const [title, setTitle] = useState('');
    const [url, setUrl] = useState('');
    const navigate = useNavigate();

    const handleSubmit = (event) => {
        event.preventDefault();
        const link = { title, url };
        fetch("http://localhost:8080/links", {  // Используем относительный URL для проксирования запросов
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(link),
          }).then(response => {
            if (response.ok) {
                navigate('/')
            }
        })
            .catch(error => console.log(error));
    };

    return (
        <div className="general" id="addLink">
            <h2>Add Link</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Title"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                />
                <input
                    type="text"
                    placeholder="URL"
                    value={url}
                    onChange={(e) => setUrl(e.target.value)}
                />
                <button class="submit" type="submit">Submit</button>
            </form>
        </div>
    );
};

export default AddLink;
