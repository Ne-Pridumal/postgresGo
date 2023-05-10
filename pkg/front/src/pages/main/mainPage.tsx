import { useCreateTicket, useTestQuery } from '@features/api/main';
import { Form } from '@shared/organisms/form';
import { ChangeEvent, FormEvent, useCallback, useState } from 'react';
import { callbackify } from 'util';

export type TMainPage = {

}

type TForm = {
  destination: string,
  departure: string,
  date: string,
  price: number,
  name: string,
}

export const MainPage = ({ }: TMainPage) => {
  const { data, isLoading, isError, mutateAsync: createTicket } = useCreateTicket()
  const [showResult, setShowResult] = useState(false)
  const [formState, setFormState] = useState<TForm>({
    date: '',
    departure: '',
    destination: '',
    price: 0,
    name: '',
  });

  const submit = () => {
    createTicket(formState)
    setShowResult(!isError)
  };

  const handleChange = useCallback((event: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setFormState((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  }, []);

  return (
    <div>
      {
        showResult
          ?
          <table className="table-auto w-full">
            <thead>
              <tr>
                <th className="px-4 py-2">Name</th>
                <th className="px-4 py-2">Amount</th>
                <th className="px-4 py-2">Departure Date</th>
              </tr>
            </thead>
            <tbody>
              <tr key={data?.name}>
                <td className="border px-4 py-2">{data?.name}</td>
                <td className="border px-4 py-2">{data?.amount}</td>
                <td className="border px-4 py-2">{data?.departureDate.toString()}</td>
              </tr>
            </tbody>
          </table>
          :
          <Form
            name={formState.name}
            submit={submit}
            callback={handleChange}
            date={formState.date}
            price={formState.price}
            departure={formState.departure}
            destination={formState.destination}
          />

      }
    </div>
  );
};
