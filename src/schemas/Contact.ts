import { model, Schema } from 'mongoose';

const ContactSchema = new Schema({
  email: { type: String },
});

export default model('users', ContactSchema);
