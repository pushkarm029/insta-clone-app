import React from 'react';
import { render, screen } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import '@testing-library/jest-dom';
import Navigation from '../navigation/App';

jest.mock('react-responsive', () => ({
  useMediaQuery: jest.fn(() => true),
}));

test('renders primary navigation links on desktop', () => {
  render(<Navigation />, { wrapper: BrowserRouter });

  expect(screen.getByText('Home')).toBeInTheDocument();
  expect(screen.getByText('Search')).toBeInTheDocument();
  expect(screen.getByText('Explore')).toBeInTheDocument();
  expect(screen.getByText('Reels')).toBeInTheDocument();
  expect(screen.getByText('Messages')).toBeInTheDocument();
  expect(screen.getByText('Profile')).toBeInTheDocument();
  expect(screen.getByText('Chill-Zone')).toBeInTheDocument();
});
