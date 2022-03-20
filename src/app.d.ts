/// <reference types="@sveltejs/kit" />

declare namespace App {
	import { User } from '@prisma/client';

	interface Locals {
		user: User;
	}

	interface Platform {}

	interface Session {
		user: User;
	}

	interface Stuff {}
}
