import React from "react";

function EndForm({finValue, setFinValue}) {
    const handleChange = (e) => {
        if (e !== null && e !== undefined) {
            const value = e.target.value
            setFinValue(value)
        }
    }

    return (
        <input type="text" placeholder="Insert your final link" className="inoneline" onChange={handleChange} style={{ margin: '10px',width:'500px',height:'30px',borderRadius:'20px',padding:'10px',backgroundColor:'#000020',border:'3px solid white',color:'white' }} required />
    );
};

export default EndForm;