import { fakerJA } from '@faker-js/faker';
import { generateUlid } from './ulid';

export const FakeUser = () => {
	return {
		user_id: generateUlid(),
		preferred_user_id: fakerJA.string.alphanumeric({ length: { min: 8, max: 16 } }),
		preferred_user_name: fakerJA.person.fullName(),
		is_public: true,
		mail_address: fakerJA.internet.email()
	};
};
