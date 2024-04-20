import React from "react";

function LinkForm({linkValue, setLinkValue}) {
    const handleChange = (e) => {
        if (e !== null && e !== undefined) {
            const value = e.target.value;
            setLinkValue(value);
        }
    }

    return (
        <input type="text" placeholder="Insert your starting wikipedia title here" className="inoneline" onChange={handleChange} style={{ margin: '10px',width:'500px',height:'30px',borderRadius:'20px',padding:'10px',backgroundColor:'#200000',border:'3px solid white',color:'white' }}required />
    );
};

export default LinkForm;