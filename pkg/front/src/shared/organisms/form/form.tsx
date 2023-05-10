import { ChangeEvent } from 'react';

export type TForm = {
  date: string,
  departure: string,
  destination: string,
  price: number,
  name: string,
  callback: (event: ChangeEvent<HTMLInputElement>) => void,
  submit: () => void
}

export const Form = (formProps: TForm) => {
  return (
    <div className="max-w-md mx-auto">
      <div className="mb-4">
        <label htmlFor="destination" className="block mb-2 font-bold text-gray-700">
          Destination:
        </label>
        <input
          type="text"
          id="destination"
          name="destination"
          value={formProps.destination}
          onChange={formProps.callback}
          className="w-full px-3 py-2 leading-tight border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="mb-4">
        <label htmlFor="departure" className="block mb-2 font-bold text-gray-700">
          Departure:
        </label>
        <input
          type="text"
          id="departure"
          name="departure"
          value={formProps.departure}
          onChange={formProps.callback}
          className="w-full px-3 py-2 leading-tight border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="mb-4">
        <label htmlFor="date" className="block mb-2 font-bold text-gray-700">
          Date:
        </label>
        <input
          id="date"
          name="date"
          type="date"
          value={formProps.date}
          onChange={formProps.callback}
          className="w-full px-3 py-2 leading-tight border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="mb-4">
        <label htmlFor="price" className="block mb-2 font-bold text-gray-700">
          Price:
        </label>
        <input
          id="price"
          name="price"
          type="number"
          value={formProps.price}
          onChange={formProps.callback}
          className="w-full px-3 py-2 leading-tight border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="mb-4">
        <label htmlFor="name" className="block mb-2 font-bold text-gray-700">
          Name:
        </label>
        <input
          id="name"
          name="name"
          type="text"
          value={formProps.name}
          onChange={formProps.callback}
          className="w-full px-3 py-2 leading-tight border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="text-center">
        <button
          type="submit"
          className="px-4 py-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700 focus:outline-none focus:shadow-outline-blue"
          onClick={formProps.submit}
        >
          Submit
        </button>
      </div>
    </div>
  );
};
