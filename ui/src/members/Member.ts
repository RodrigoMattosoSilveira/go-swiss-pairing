export class Member {
    id: string | undefined;
    first: string = '';
    email: string = '';
    isActive: boolean = false

    constructor(initializer?: any) {
        if (!initializer) return;
        if (initializer.id) this.id = initializer.id;
        if (initializer.first) this.first = initializer.first;
        if (initializer.email) this.email = initializer.email;
        if (initializer.isActive) this.isActive = initializer.isActive;
    }
}