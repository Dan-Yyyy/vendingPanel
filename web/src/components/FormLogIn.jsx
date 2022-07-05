import React from "react";
import styled from "styled-components";
import { useDispatch } from 'react-redux';
import { useFormik } from 'formik';
import { Input } from "./Input";
import { Button } from "./Button";
import { Error } from "./Error";
import { authThunk } from './../redux/auth/authThunk';

const InputContainer = styled.div`
  position: relative;
  margin-bottom: 48px;
`;

export const FormLogIn = () => {
  const dispatch = useDispatch();
  const validate = (values) => {
    const errors = {};
  
      if (!values.email) {
        errors.email = "Поле должно быть заполнено";
      } else if (!/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(values.email)) {
        errors.email = "Не корректно введен email";
      }
      if (!values.password) {
        errors.password = "Поле должно быть заполнено";
      } 
      return errors;
  };

  const formik = useFormik({
    initialValues: {
      email: '',
      password: '',
    },
    validate,
    onSubmit: values => {
      dispatch(authThunk(JSON.stringify(values, null, 2)));
    },
  });

  return(
    <form onSubmit={formik.handleSubmit}>
      <InputContainer>
        <Input name={'email'} type={'email'}
          placeholder="email"
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          value={formik.values.email}/>
        {formik.touched.email && formik.errors.email ? <Error>{formik.errors.email}</Error> : null}
      </InputContainer>
      <InputContainer>
        <Input name={'password'} type={'password'}
          placeholder="password"
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          value={formik.values.password}/>
        {formik.touched.password && formik.errors.password ? <Error>{formik.errors.password}</Error> : null}
      </InputContainer>
      <Button type="submit" disabled={!formik.isValid} color={'#A0694B'}>Войти</Button>
    </form>
  );
}