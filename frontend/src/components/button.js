import React from "react";
import Button from "react-bootstrap/Button";
import axios from "axios";

const URL = "http://localhost:4000/newmeeting";
const tokenURL = "http://localhost:4000/token";

function LargeButton() {
  const [post, setPost] = React.useState(null);
  function createPost() {
    axios.post(URL).then((response) => {
      console.log(response);
    });
  }

  return (
    <Button onClick={createPost} variant="primary" size="lg">
      Create a Meeting
    </Button>
  );
}
export default LargeButton;
