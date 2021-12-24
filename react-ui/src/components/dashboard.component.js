import React from 'react';
import Login from './login.component';
import {Col, Container, Row} from 'react-bootstrap';


export default function Dashboard({setToken}) {

    if (!sessionStorage.getItem('token')) {
        return <Login setToken={setToken}/>
    }
    return (
        <Container fluid className="dashboard-container" style={{display: 'block', backgroundColor: 'yellow'}}>
            <Row className="h-50">
                <Col className="dashboard-col">
                    <h3>Col 1</h3>
                </Col>
                <Col className="dashboard-col">
                    <h3>Col 2</h3>
                </Col>
            </Row>
            <Row className="h-50">
                <Col className="dashboard-col">
                    <h3>Row 3</h3>
                </Col>
            </Row>
        </Container>
    );
}