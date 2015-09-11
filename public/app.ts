//Typescript
/// <reference path="typings/angular2/angular2.d.ts" />
import {Component,View, bootstrap, NgFor} from 'angular2/angular2';

@Component({
  selector: 'user-input'
})

@View({
  directives: [NgFor],
  templateUrl : 'user-input.html',
})

// Component controller
class UserComponent {
  username: string;

  constructor() {},

  selectUsername(name: string) {
    this.username = name;

    // TODO: Socket.io emit
  }
}

bootstrap(UserComponent);
