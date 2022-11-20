import { Member } from './Member';
const baseUrl = 'http://localhost:4000';
const url = `${baseUrl}/members`;

function translateStatusToErrorMessage(status: number) {
    switch (status) {
        case 401:
            return 'Please login again.';
        case 403:
            return 'You do not have permission to view the member(s).';
        default:
            return 'There was an error retrieving the member(s). Please try again.';
    }
}

function checkStatus(response: any) {
    if (response.ok) {
        return response;
    } else {
        const httpErrorInfo = {
            status: response.status,
            statusText: response.statusText,
            url: response.url,
        };
        console.log(`log server http error: ${JSON.stringify(httpErrorInfo)}`);

        let errorMessage = translateStatusToErrorMessage(httpErrorInfo.status);
        throw new Error(errorMessage);
    }
}

function parseJSON(response: Response) {
    return response.json();
}

// eslint-disable-next-line
function delay(ms: number) {
    return function (x: any): Promise<any> {
        return new Promise((resolve) => setTimeout(() => resolve(x), ms));
    };
}

function convertToMemberModels(data: any[]): Member[] {
    let members: Member[] = data.map(convertToMemberModel);
    return members;
}

function convertToMemberModel(item: any): Member {
    return new Member(item);
}

const memberAPI = {
    get(page = 1, limit = 20) {
        return fetch(`${url}?_page=${page}&_limit=${limit}&_sort=first`)
            // .then(delay(3000))
            .then(checkStatus)
            .then(parseJSON)
            .then(convertToMemberModels)
            .catch((error: TypeError) => {
                console.log('log client error ' + error);
                throw new Error(
                    'There was an error retrieving the members. Please try again.'
                );
            });
    },
};

export { memberAPI };