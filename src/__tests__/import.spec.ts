import mongoose from 'mongoose';
import Contact from '@schemas/Contact';

describe('When import a csv file', () => {
  beforeAll(async () => {
    if (!process.env.MONGO_URL) {
      throw new Error('MongoDB server not initialized');
    }

    await mongoose.connect(process.env.MONGO_URL, {
      useNewUrlParser: true,
      useUnifiedTopology: true,
      useCreateIndex: true,
    });
  });

  afterAll(async () => {
    mongoose.connection.close();
  });

  beforeEach(async () => {
    await Contact.deleteMany({});
  });

  it('shold be able to import a new contacts', async () => {
    await Contact.create({ email: 'tassiorego@gmail.com' });

    const contacts = await Contact.find();

    expect(contacts).toEqual(
      expect.arrayContaining([
        expect.objectContaining({
          email: 'tassiorego@gmail.com',
        }),
      ]),
    );
  });
});
