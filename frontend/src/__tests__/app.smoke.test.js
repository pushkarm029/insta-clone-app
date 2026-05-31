import React from 'react';
import { render, screen } from '@testing-library/react';
import App from '../App';

jest.mock('../firebase', () => ({
  __esModule: true,
  default: {},
}));

jest.mock('../components/auth/SignIn', () => function MockSignIn() {
  return <div>Sign in screen</div>;
});

jest.mock('../navigator', () => function MockNavigator() {
  return <div>App navigation</div>;
});

test('renders the signed-out app shell', () => {
  render(<App />);

  expect(screen.getByText('Sign in screen')).toBeInTheDocument();
});
