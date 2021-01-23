import React from "react";
import {Form, Input, Checkbox, Button} from "antd";
import "./Login.css"
import axios from "axios";

export default class Login extends React.Component<any, any> {
    layout = {
        labelCol: { span: 8 },
        wrapperCol: { span: 16 },
    };
    tailLayout = {
        wrapperCol: { offset: 8, span: 16 },
    };
    onFinish = (values: any) => {
        axios.post("/api/user/login", {
            username: values['username'],
            password: values['password']
        }).then(function (response) {

        }).catch(function (error) {
            alert(error)
        })
    };
    onFinishFailed = (errorInfo : any) => {
        console.log('Failed:', errorInfo);
    };

    render() {
        return (
            <Form
                {...this.layout}
                name="basic"
                initialValues={{ remember: true }}
                onFinish={this.onFinish}
                onFinishFailed={this.onFinishFailed}
            >
                <Form.Item
                    label="Username"
                    name="username"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Password"
                    name="password"
                    rules={[{ required: true, message: 'Please input your password!' }]}
                >
                    <Input.Password />
                </Form.Item>

                <Form.Item {...this.tailLayout} name="remember" valuePropName="checked">
                    <Checkbox>Remember me</Checkbox>
                </Form.Item>

                <Form.Item {...this.tailLayout}>
                    <Button type="primary" htmlType="submit">
                        Submit
                    </Button>
                </Form.Item>
            </Form>
        )
    }
}