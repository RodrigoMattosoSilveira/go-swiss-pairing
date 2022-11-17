export class Member {
    id: string | undefined;
    first: string = '';
    last: string = '';
    email: string = '';
    password: string = '';
    cell: string = '';
    rating: number = 1200;
    isActive: boolean = false
    imageUrl: string = 'picture_placeholder.jpeg';

    constructor(initializer?: any) {
        if (!initializer) return;
        if (initializer.id) this.id = initializer.id;
        if (initializer.first) this.first = initializer.first;
        if (initializer.last) this.last = initializer.last;
        if (initializer.email) this.email = initializer.email;
        if (initializer.password) this.password = initializer.password;
        if (initializer.cell) this.cell = initializer.cell;
        if (initializer.rating) this.rating = initializer.rating;
        if (initializer.isActive) this.isActive = initializer.isActive;
        if (initializer.imageUrl) this.imageUrl = initializer.imageUrl;
    }
}
