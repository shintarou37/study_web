import { useForm, SubmitHandler } from "react-hook-form"

const onSubmit = (data: any) => {
  console.log(data)
}
function Form() {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm();
  return (
    <div>
      <form onSubmit={handleSubmit(onSubmit)}>
        <label htmlFor="name">Name</label>
        <input id="name" type="text" {...register("name", { required: "name is required.", minLength: { value: 4, message: "Name must be at latest 4 characters long." } })}></input>
        <p>{errors.name?.message as React.ReactNode}</p>
        <button type="submit">Submit</button>
      </form>
    </div>
  )
}

export default Form