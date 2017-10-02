export class ProjectModel {

  // This project definition is based on CINCO project class
  id: string;
  name: string;
  description: string;
  managers: string;
  members: any;
  status: string;
  category: string;
  sector: string;
  url: string;
  startDate: string;
  agreementRef: string;
  mailingListType: string;
  emailAliasType: string;
  address: {
    address: {
      administrativeArea: string,
      country: string,
      localityName: string,
      postalCode: string,
      thoroughfare: string
    },
    type: string
  }
  config: {
    logoRef: string
  }

  constructor() {
  }

}
