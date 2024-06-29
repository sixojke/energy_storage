import { Component, inject, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { TableComponent } from './common-ui/table/table.component';
import { HydrogenStorageService } from './data/services/hydrogen-storage.service';
import { HydrogenStorage } from './data/interfaces/hydrogen-storage.interface';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, TableComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  hydrogenStorageService = inject(HydrogenStorageService)
  hydrogenStorageData: HydrogenStorage[] = []

  constructor() {
    this.hydrogenStorageService.getData().subscribe(val => {
      console.log(val)
      this.hydrogenStorageData = val
      console.log(this.hydrogenStorageData)
    })
  }
}
