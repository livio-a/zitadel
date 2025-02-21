import { SelectionModel } from '@angular/cdk/collections';
import { Component, EventEmitter, Input, OnInit, Output, ViewChild } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatTableDataSource } from '@angular/material/table';
import { TranslateService } from '@ngx-translate/core';
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';
import { Moment } from 'moment';
import { BehaviorSubject, Observable } from 'rxjs';
import { AddKeyDialogComponent, AddKeyDialogType } from 'src/app/modules/add-key-dialog/add-key-dialog.component';
import { ShowKeyDialogComponent } from 'src/app/modules/show-key-dialog/show-key-dialog.component';
import { Key, KeyType } from 'src/app/proto/generated/zitadel/auth_n_key_pb';
import { ListAppKeysResponse } from 'src/app/proto/generated/zitadel/management_pb';
import { ManagementService } from 'src/app/services/mgmt.service';
import { ToastService } from 'src/app/services/toast.service';

import { PageEvent, PaginatorComponent } from '../paginator/paginator.component';

@Component({
  selector: 'app-client-keys',
  templateUrl: './client-keys.component.html',
  styleUrls: ['./client-keys.component.scss'],
})
export class ClientKeysComponent implements OnInit {
  @Input() projectId!: string;
  @Input() appId!: string;

  @ViewChild(PaginatorComponent) public paginator!: PaginatorComponent;
  public dataSource: MatTableDataSource<Key.AsObject> = new MatTableDataSource<Key.AsObject>();
  public selection: SelectionModel<Key.AsObject> = new SelectionModel<Key.AsObject>(true, []);
  public keyResult!: ListAppKeysResponse.AsObject;
  private loadingSubject: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
  public loading$: Observable<boolean> = this.loadingSubject.asObservable();
  @Input() public displayedColumns: string[] = ['select', 'id', 'type', 'creationDate', 'expirationDate', 'actions'];

  @Output() public changedSelection: EventEmitter<Array<Key.AsObject>> = new EventEmitter();

  constructor(public translate: TranslateService, private mgmtService: ManagementService, private dialog: MatDialog,
    private toast: ToastService) {
    this.selection.changed.subscribe(() => {
      this.changedSelection.emit(this.selection.selected);
    });
  }

  public ngOnInit(): void {
    this.getData(10, 0);
  }


  public isAllSelected(): boolean {
    const numSelected = this.selection.selected.length;
    const numRows = this.dataSource.data.length;
    return numSelected === numRows;
  }

  public masterToggle(): void {
    this.isAllSelected() ?
      this.selection.clear() :
      this.dataSource.data.forEach(row => this.selection.select(row));
  }


  public changePage(event: PageEvent): void {
    this.getData(event.pageSize, event.pageIndex * event.pageSize);
  }

  public deleteKey(key: Key.AsObject): void {
    this.mgmtService.removeAppKey(this.projectId, this.appId, key.id).then(() => {
      this.selection.clear();
      this.toast.showInfo('USER.TOAST.SELECTEDKEYSDELETED', true);
      this.getData(10, 0);
    }).catch(error => {
      this.toast.showError(error);
    });
  }

  public openAddKey(): void {
    const dialogRef = this.dialog.open(AddKeyDialogComponent, {
      data: {},
      width: '400px',
    });

    dialogRef.afterClosed().subscribe(resp => {
      if (resp) {
        const type: KeyType = resp.type;

        let date: Timestamp | undefined;

        if (resp.date as Moment) {
          const ts = new Timestamp();
          const milliseconds = resp.date.toDate().getTime();
          const seconds = Math.abs(milliseconds / 1000);
          const nanos = (milliseconds - seconds * 1000) * 1000 * 1000;
          ts.setSeconds(seconds);
          ts.setNanos(nanos);
          date = ts;
        }

        if (type) {
          return this.mgmtService.addAppKey(
            this.projectId,
            this.appId,
            type,
            date ? date : undefined,
          ).then((response) => {
            if (response) {
              setTimeout(() => {
                this.refreshPage();
              }, 1000);

              this.dialog.open(ShowKeyDialogComponent, {
                data: {
                  key: response,
                  type: AddKeyDialogType.AUTHNKEY,
                },
                width: '400px',
              });
            }
          }).catch((error: any) => {
            this.toast.showError(error);
          });
        }
      }
    });
  }

  private async getData(limit: number, offset: number): Promise<void> {
    this.loadingSubject.next(true);
    if (this.projectId && this.appId) {
      this.mgmtService.listAppKeys(this.projectId, this.appId, limit, offset).then(resp => {
        this.keyResult = resp;
        this.dataSource.data = this.keyResult.resultList;
        this.loadingSubject.next(false);
      }).catch((error: any) => {
        this.toast.showError(error);
        this.loadingSubject.next(false);
      });
    }
  }

  public refreshPage(): void {
    this.getData(this.paginator.pageSize, this.paginator.pageIndex * this.paginator.pageSize);
  }
}
