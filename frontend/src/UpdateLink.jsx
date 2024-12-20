import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';

const UpdateLink = () => {
    const { id } = useParams();
    const [link, setLink] = useState({ title: '', url: '' });
    const navigate = useNavigate();

    useEffect(() => {
        fetch(`http://localhost:8080/links/${id}`)
            .then((response) => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then((data) => setLink(data))
            .catch((error) => console.log(error));
    }, [id]);

    const handleChange = (event) => {
        const { name, value } = event.target;
        setLink((prevLink) => ({
            ...prevLink,
            [name]: value,
        }));
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        fetch(`http://localhost:8080/links/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(link),
        })
            .then((response) => {
                if (response.ok) {
                    navigate('/');
                }
            })
            .catch((error) => console.log(error));
    };

    return (
        <div className="general" id="updateLink">
            <h2>Update Link</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    name="title"
                    placeholder="Title"
                    value={link.title}
                    onChange={handleChange}
                />
                <input
                    type="text"
                    name="url"
                    placeholder="URL"
                    value={link.url}
                    onChange={handleChange}
                />
                <button className="submit" type="submit">Update</button>
            </form>
        </div>
    );
};

export default UpdateLink;
