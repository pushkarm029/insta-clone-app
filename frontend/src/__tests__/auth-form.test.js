import React from 'react';
import { render, screen } from '@testing-library/react';
import { Provider } from 'react-redux';
import store from '../components/redux/store';
import Login from '../components/auth/LoginAccountForm';
import CreateAccountForm from '../components/auth/CreateAccountForm';

jest.mock('firebase/auth', () => ({
  getAuth: jest.fn(() => ({})),
  signInWithEmailAndPassword: jest.fn(),
  createUserWithEmailAndPassword: jest.fn(),
}));

jest.mock('firebase/firestore', () => ({
  getFirestore: jest.fn(() => ({})),
  collection: jest.fn(),
  addDoc: jest.fn(),
}));

test('renders login email and password inputs', () => {
  render(
    <Provider store={store}>
      <Login onCreateForm={jest.fn()} onLogin={jest.fn()} />
    </Provider>
  );

  expect(screen.getByPlaceholderText('Enter your email')).toBeInTheDocument();
  expect(screen.getByPlaceholderText('Password')).toBeInTheDocument();
});

test('renders create account required fields', () => {
  render(<CreateAccountForm onCreateForm={jest.fn()} />);

  expect(screen.getByPlaceholderText('Full Name (Required)')).toBeInTheDocument();
  expect(screen.getByPlaceholderText('Username (Required)')).toBeInTheDocument();
  expect(screen.getByPlaceholderText('Email (Required)')).toBeInTheDocument();
  expect(screen.getByPlaceholderText('Password (Required)')).toBeInTheDocument();
});
