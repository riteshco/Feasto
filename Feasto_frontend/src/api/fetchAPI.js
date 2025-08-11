export const GetProducts = async () => {
    try {
        const response = await fetch(
            "http://localhost:3000/api/all-products"
        );
        products = response.json()
    }
    catch (error) {
        console.log(error);
    }
}

export const GetUsers = async () => {
    try {
        const response = await fetch(
            "http://localhost:3000/api/user"
        );
        users = response.json()
    }
    catch (error) {
        console.log(error);
    }
}

export const GetUserByID = async (id) => {
    try {
        const response = await fetch(
            `http://localhost:3000/api/user/${id}`
        );
        user = response.json()
    }
    catch (error) {
        console.log(error);
    }
}