import {
  Datastore as GoogleDatastore,
  DatastoreOptions,
  Key,
} from '@google-cloud/datastore';

export class Datastore {
  private client: GoogleDatastore;
  private kind: string;

  constructor() {
    const options: DatastoreOptions = {};

    this.client = new GoogleDatastore(options);
    this.kind = 'cateiru.com';
  }

  private createKey(name: string[]): Key {
    return this.client.key([this.kind, ...name]);
  }

  public async set<T>(keyName: string, obj: T) {
    const entity = {
      Key: this.createKey([keyName]),
      data: obj,
    };

    await this.client.save(entity);
  }

  public async get<T>(keyName: string): Promise<T[]> {
    const [entities] = await this.client.get(this.createKey([keyName]));

    if (Array.isArray(entities)) {
      return entities as T[];
    } else {
      return [entities] as T[];
    }
  }

  public async delete(keyName: string) {
    await this.client.delete(this.createKey([keyName]));
  }
}
