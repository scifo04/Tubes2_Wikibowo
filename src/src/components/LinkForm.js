import React, { useState } from "react";

function LinkForm({linkValue, setLinkValue}) {
    const [suggestions,setSuggestions] = useState([])
    const handleChange = async (e) => {
        if (e !== null && e !== undefined) {
            const value = e.target.value;
            setLinkValue(value);
            try {
                const response = await fetch(
                    `https://api.allorigins.win/get?url=${encodeURIComponent(
                      `https://en.wikipedia.org/w/api.php?action=query&format=json&list=search&srsearch=${value}`
                    )}`
                  );
                const data = await response.json()
                const jsonData = JSON.parse(data.contents);
                // console.log(data)
                // console.log(jsonData)
            
                // console.log(jsonData.query.search)
                console.log(jsonData.query.search.map(item => item.title))
                setSuggestions(jsonData.query.search.map(item => item.title))

            } catch (error) {
                console.error('Error fetching suggestions:', error);
                setSuggestions([]);
            }
        }
    }
    return (
        <div>
            <input 
                type="text" 
                placeholder="Insert your starting wikipedia title here"
                className="inoneline" 
                onChange={handleChange} 
                style={{ 
                    margin: '10px',
                    width:'500px',
                    height:'30px',
                    borderRadius:'20px',
                    padding:'10px',
                    backgroundColor:'#200000',
                    border:'3px solid white',
                    color:'white' }}
                    required 
            />
            <ul>
                {suggestions.map((title,index) => (
                    <li key={index}>{title}</li>
                ))}
            </ul>
        </div>
    );
};

export default LinkForm;