import { firebaseConfig } from '../config/firebaseConfig';

test('loads required firebase web config values', () => {
  expect(firebaseConfig.apiKey).toBeTruthy();
  expect(firebaseConfig.authDomain).toContain('firebaseapp.com');
  expect(firebaseConfig.projectId).toBeTruthy();
  expect(firebaseConfig.storageBucket).toContain('appspot.com');
  expect(firebaseConfig.appId).toBeTruthy();
});
